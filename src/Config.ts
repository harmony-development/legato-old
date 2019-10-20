import { readFileSync } from 'fs';

interface IConfig {
  jwtsecret: string;
}

export class Config {
  config: IConfig;

  constructor() {
    const readconfig = readFileSync('config.json', 'utf8');
    this.config = JSON.parse(readconfig);
  }
}
