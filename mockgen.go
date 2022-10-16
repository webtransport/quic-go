package quic

//go:generate sh -c "./mockgen_private.sh quic mock_send_conn_test.go github.com/btwiuse/quic sendConn"
//go:generate sh -c "./mockgen_private.sh quic mock_sender_test.go github.com/btwiuse/quic sender"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_internal_test.go github.com/btwiuse/quic streamI"
//go:generate sh -c "./mockgen_private.sh quic mock_crypto_stream_test.go github.com/btwiuse/quic cryptoStream"
//go:generate sh -c "./mockgen_private.sh quic mock_receive_stream_internal_test.go github.com/btwiuse/quic receiveStreamI"
//go:generate sh -c "./mockgen_private.sh quic mock_send_stream_internal_test.go github.com/btwiuse/quic sendStreamI"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_sender_test.go github.com/btwiuse/quic streamSender"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_getter_test.go github.com/btwiuse/quic streamGetter"
//go:generate sh -c "./mockgen_private.sh quic mock_crypto_data_handler_test.go github.com/btwiuse/quic cryptoDataHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_frame_source_test.go github.com/btwiuse/quic frameSource"
//go:generate sh -c "./mockgen_private.sh quic mock_ack_frame_source_test.go github.com/btwiuse/quic ackFrameSource"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_manager_test.go github.com/btwiuse/quic streamManager"
//go:generate sh -c "./mockgen_private.sh quic mock_sealing_manager_test.go github.com/btwiuse/quic sealingManager"
//go:generate sh -c "./mockgen_private.sh quic mock_unpacker_test.go github.com/btwiuse/quic unpacker"
//go:generate sh -c "./mockgen_private.sh quic mock_packer_test.go github.com/btwiuse/quic packer"
//go:generate sh -c "./mockgen_private.sh quic mock_mtu_discoverer_test.go github.com/btwiuse/quic mtuDiscoverer"
//go:generate sh -c "./mockgen_private.sh quic mock_conn_runner_test.go github.com/btwiuse/quic connRunner"
//go:generate sh -c "./mockgen_private.sh quic mock_quic_conn_test.go github.com/btwiuse/quic quicConn"
//go:generate sh -c "./mockgen_private.sh quic mock_packet_handler_test.go github.com/btwiuse/quic packetHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_unknown_packet_handler_test.go github.com/btwiuse/quic unknownPacketHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_packet_handler_manager_test.go github.com/btwiuse/quic packetHandlerManager"
//go:generate sh -c "./mockgen_private.sh quic mock_multiplexer_test.go github.com/btwiuse/quic multiplexer"
//go:generate sh -c "./mockgen_private.sh quic mock_batch_conn_test.go github.com/btwiuse/quic batchConn"
//go:generate sh -c "mockgen -package quic -self_package github.com/btwiuse/quic -destination mock_token_store_test.go github.com/btwiuse/quic TokenStore"
//go:generate sh -c "mockgen -package quic -self_package github.com/btwiuse/quic -destination mock_packetconn_test.go net PacketConn"
