import { $ } from "zx";
import { start } from "./checkBalance";

$.verbose = false;

test("checks balance with no arguments", async () => {
  await expect(start()).rejects.toThrowError("chain not defined");
});

test("checks balance with chain set to gridchain", async () => {
  process.argv = ["node", "jest", "--chain", "gridchain"];

  await expect(start()).rejects.toThrowError(
    "missing requirement argument: --name"
  );
});

test("checks balance with chain set to gridchain and name set to XXX", async () => {
  process.argv = ["node", "jest", "--chain", "gridchain", "--name", "XXX"];

  await expect(start()).rejects.toThrowError(
    "Error: decoding bech32 failed: invalid bech32 string length 3"
  );
});

test("checks balance with chain set to gridchain and network set to testnet-1", async () => {
  process.argv = [
    "node",
    "jest",
    "--chain",
    "gridchain",
    "--network",
    "testnet-1",
  ];

  await expect(start()).rejects.toThrowError(
    "missing requirement argument: --name"
  );
});
