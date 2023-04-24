import { $, nothrow } from "zx";
import { send } from "../lib/send.mjs";

export async function createRelayer({
  gridChainProps,
  otherChainProps,
  registryFrom = `/tmp/localnet/config/registry`,
}) {
  const { chain, home } = otherChainProps;
  const relayerHome = `${home}/relayer`;

  await nothrow($`mkdir -p ${relayerHome}`);
  await nothrow(
    $`ibc-setup init --home ${relayerHome} --registry-from ${registryFrom} --src ${gridChainProps.chain} --dest ${chain}`
  );

  let addresses = await $`ibc-setup keys list --home ${relayerHome}`;
  addresses = addresses.toString().split("\n");

  const gridChainAddress = addresses
    .find((item) => item.includes(`${gridChainProps.chain}`))
    .replace(`${gridChainProps.chain}: `, ``);
  const otherChainAddress = addresses
    .find((item) => item.includes(`${chain}`))
    .replace(`${chain}: `, ``);

  console.log(`gridChainAddress: ${gridChainAddress}`);
  console.log(`otherChainAddress: ${otherChainAddress}`);

  await send({
    ...otherChainProps,
    src: `${chain}-source`,
    dst: otherChainAddress,
    amount: 10e10,
    node: `tcp://127.0.0.1:${otherChainProps.rpcPort}`,
  });

  return {
    gridChainAddress,
    otherChainAddress,
    gridSendRequest: {
      ...gridChainProps,
      src: `${gridChainProps.chain}-source`,
      dst: gridChainAddress,
      amount: 10e10,
      node: `tcp://127.0.0.1:${gridChainProps.rpcPort}`,
    },
  };
}
