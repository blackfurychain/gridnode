import { $ } from "zx";
import { start } from "./checkBalance";

$.verbose = false;

test("checks balance with no arguments", async () => {
  await expect(start()).rejects.toThrowError("chain not defined");
});

test("checks balance with chain set to gridironchain", async () => {
  process.argv = ["node", "jest", "--chain", "gridironchain"];

  await expect(start()).rejects.toThrowError(
    "missing requirement argument: --name"
  );
});

test("checks balance with chain set to gridironchain and name set to XXX", async () => {
  process.argv = ["node", "jest", "--chain", "gridironchain", "--name", "XXX"];

  await expect(start()).rejects.toThrowError(
    "Error: decoding bech32 failed: invalid bech32 string length 3"
  );
});

test("checks balance with chain set to gridironchain and network set to testnet-1", async () => {
  process.argv = [
    "node",
    "jest",
    "--chain",
    "gridironchain",
    "--network",
    "testnet-1",
  ];

  await expect(start()).rejects.toThrowError(
    "missing requirement argument: --name"
  );
});
