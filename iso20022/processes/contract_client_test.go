package processes_test

import (
	"context"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	coreumlogger "github.com/CoreumFoundation/coreum-tools/pkg/logger"
	coreumchainclient "github.com/CoreumFoundation/coreum/v4/pkg/client"
	coreumchainconstant "github.com/CoreumFoundation/coreum/v4/pkg/config/constant"
	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
	"github.com/CoreumFoundation/iso20022-client/iso20022/compress"
	"github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
	isoqueue "github.com/CoreumFoundation/iso20022-client/iso20022/queue"
)

func TestMain(m *testing.M) {
	coreum.SetSDKConfig(coreumchainconstant.AddressPrefixDev)
	m.Run()
}

func TestContractClient_Start(t *testing.T) {
	tests := []struct {
		name                  string
		contractClientBuilder func(ctrl *gomock.Controller) processes.ContractClient
		addressBookBuilder    func(ctrl *gomock.Controller) processes.AddressBook
		cryptographyBuilder   func(ctrl *gomock.Controller) processes.Cryptography
		parserBuilder         func(ctrl *gomock.Controller) processes.Parser
		messageQueueBuilder   func(ctrl *gomock.Controller, queue chan [][]byte) processes.MessageQueue
		dtifBuilder           func(ctrl *gomock.Controller) processes.Dtif
		run                   func(messageQueue processes.MessageQueue) error
	}{
		{
			name: "sending_one_message",
			contractClientBuilder: func(ctrl *gomock.Controller) processes.ContractClient {
				contractClientMock := NewMockContractClient(ctrl)
				contractClientMock.EXPECT().IsInitialized().Return(true)
				contractClientMock.EXPECT().GetNewMessages(gomock.Any(), gomock.Any(), gomock.Any()).Return([]coreum.Message{}, nil).AnyTimes()
				contractClientMock.EXPECT().GetActiveSessions(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]coreum.Session{}, nil).AnyTimes()
				contractClientMock.EXPECT().IssueNFTClass(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
				contractClientMock.EXPECT().BroadcastMessages(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
				contractClientMock.EXPECT().SendMessages(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
				return contractClientMock
			},
			addressBookBuilder: func(ctrl *gomock.Controller) processes.AddressBook {
				addressBookMock := NewMockAddressBook(ctrl)
				addressBookMock.EXPECT().Lookup(gomock.Any()).Return(&addressbook.Address{
					Bech32EncodedAddress: "devcore1kdujjkp8u0j9lww3n7qs7r5fwkljelvecsq43d",
					PublicKey:            "A2nYC1ZLFxVLL3kyGUGF4Hjlpwsd+FS7jmxIWahM0P5V",
					Party: addressbook.Party{
						Identification: addressbook.Identification{
							BusinessIdentifiersCode: "6P9YGUDF",
						},
					},
				}, true).Times(2)
				return addressBookMock
			},
			cryptographyBuilder: func(ctrl *gomock.Controller) processes.Cryptography {
				cryptographyMock := NewMockCryptography(ctrl)
				cryptographyMock.EXPECT().GenerateSharedKeyByPrivateKeyName(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte("Thirty-two bytes long shared key"), nil)
				cryptographyMock.EXPECT().EncryptSymmetric(gomock.Any(), gomock.Any()).Return([]byte("encrypted"))
				return cryptographyMock
			},
			parserBuilder: func(ctrl *gomock.Controller) processes.Parser {
				parserMock := NewMockParser(ctrl)
				parserMock.EXPECT().ExtractMessageAndMetadataFromIsoMessage(gomock.Any()).Return(
					nil,
					processes.Metadata{
						ID: "abc123",
						Sender: &addressbook.Party{
							Identification: addressbook.Identification{
								BusinessIdentifiersCode: "B61NZT4Y",
							},
						},
						Receiver: &addressbook.Party{
							Identification: addressbook.Identification{
								BusinessIdentifiersCode: "6P9YGUDF",
							},
						},
					},
					nil,
					nil,
					nil,
				)
				parserMock.EXPECT().GetSupplementaryDataWithCorrectClearingSystem(gomock.Any(), gomock.Any()).Return(nil, false)
				return parserMock
			},
			messageQueueBuilder: func(ctrl *gomock.Controller, queue chan [][]byte) processes.MessageQueue {
				queueMock := NewMockMessageQueue(ctrl)
				queueMock.EXPECT().PushToSend(gomock.Any(), gomock.Any())
				queueMock.EXPECT().PopFromSend(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, n int, dur time.Duration) [][]byte {
						return <-queue
					},
				).MinTimes(1).MaxTimes(2)
				queueMock.EXPECT().SetStatus("abc123", isoqueue.StatusSent)
				queueMock.EXPECT().Close().DoAndReturn(func() {
					close(queue)
				})
				queue <- [][]byte{[]byte("hello world")}
				return queueMock
			},
			run: func(messageQueue processes.MessageQueue) error {
				messageQueue.PushToSend("id", []byte("hello world"))
				return nil
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := coreumlogger.WithLogger(context.Background(), coreumlogger.New(coreumlogger.ToolDefaultConfig))
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

			ctrl := gomock.NewController(t)
			logMock := logger.NewAnyLogMock(ctrl)

			var contractClient processes.ContractClient
			if tt.contractClientBuilder != nil {
				contractClient = tt.contractClientBuilder(ctrl)
			}
			var addressBook processes.AddressBook
			if tt.addressBookBuilder != nil {
				addressBook = tt.addressBookBuilder(ctrl)
			}
			var cryptography processes.Cryptography
			if tt.cryptographyBuilder != nil {
				cryptography = tt.cryptographyBuilder(ctrl)
			}
			var parser processes.Parser
			if tt.parserBuilder != nil {
				parser = tt.parserBuilder(ctrl)
			}
			var messageQueue processes.MessageQueue
			if tt.messageQueueBuilder != nil {
				queue := make(chan [][]byte, 1)
				messageQueue = tt.messageQueueBuilder(ctrl, queue)
			}
			var dtif processes.Dtif
			if tt.dtifBuilder != nil {
				dtif = tt.dtifBuilder(ctrl)
			}
			t.Cleanup(cancel)
			go func() {
				go func() {
					runRrr := tt.run(messageQueue)
					require.NoError(t, runRrr)
				}()
				<-ctx.Done()
				messageQueue.Close()
			}()

			cfg := processes.ContractClientProcessConfig{
				CoreumContractAddress: genAccount(),
				ClientAddress:         genAccount(),
				ClientKeyName:         "abc",
				PollInterval:          time.Second,
			}
			comp, err := compress.New()
			require.NoError(t, err)

			client, err := processes.NewContractClientProcess(cfg, logMock, comp, coreumchainclient.Context{}, addressBook, contractClient, cryptography, parser, messageQueue, dtif)
			require.NoError(t, err)

			err = client.Start(ctx)
			if err == nil || !errors.Is(err, context.DeadlineExceeded) {
				require.NoError(t, err)
			}
		})
	}
}

func genAccount() sdk.AccAddress {
	return sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
}
