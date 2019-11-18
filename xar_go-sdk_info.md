This project extends the original binance-chain go-sdk with methods and structures suitable for interactions with XAR. For now, it extends the original DEX client with logic for rest-api interactions with a node of XAR api.

**How to use it?**   
If your project already includes the original go-sdk, you just need to replace it with this project. You'll become able to use new client features, and your previous interactions with go-sdk wont be affected. Alternatively, nothing but import requered.

**How to test special xar logic?**   
All the logic for a special interactions is added to `TransactionClient` and `QueryClient` interfaces below the original go-sdk methods. You can see the examples of usage inside the following folder: `example/xar_api_test`.   
Basically you just need to create a new Dex client with `NewCustomClient()` method, and call the method you need.  
