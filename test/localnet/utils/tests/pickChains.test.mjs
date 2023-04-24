import { pickChains } from "../pickChains.mjs";

test("pick chains", () => {
  const result = pickChains({ chain: "gridnode,cosmos,akash" });

  expect(result).toMatchSnapshot();
});
