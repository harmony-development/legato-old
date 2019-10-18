"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var chalk_1 = __importDefault(require("chalk"));
var Server_1 = require("./Server");
exports.harmonyServer = new Server_1.Server(4000);
exports.harmonyServer.open().then(function () {
    console.log(chalk_1.default.green('Successfully listening on port 4000'));
});
