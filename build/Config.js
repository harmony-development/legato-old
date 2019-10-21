"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var fs_1 = require("fs");
var Config = /** @class */ (function () {
    function Config() {
        var readconfig = fs_1.readFileSync('config.json', 'utf8');
        this.config = JSON.parse(readconfig);
    }
    return Config;
}());
exports.Config = Config;
