FROM ethereum/client-go:v1.13.14

EXPOSE 8545
EXPOSE 9545

COPY gethData /

ENTRYPOINT ["geth", "--rpc.allow-unprotected-txs", "--http", "--http.addr", "0.0.0.0", "--http.corsdomain", "*", "--http.vhosts", "*", "--ws", "--ws.origins", "*", "--ws.addr", "0.0.0.0", "--dev", "--dev.period", "1", "--datadir", "/geth_data"]