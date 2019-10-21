"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var chalk_1 = __importDefault(require("chalk"));
var Server_1 = require("./Server");
var Config_1 = require("./Config");
var PORT = 4000;
exports.harmonyServer = new Server_1.Server(PORT);
exports.config = new Config_1.Config();
exports.harmonyServer.open().then(function () {
    console.log(chalk_1.default.green("Listening on port " + PORT));
});
