import { chains } from 'chain-registry';

const chainNames = [ 'cosmoshub'];

export const chainOptions = chainNames.map(
  (chainName) => chains.find((chain) => chain.chain_name === chainName)!
);