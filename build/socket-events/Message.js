"use strict";
var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var lodash_1 = __importDefault(require("lodash"));
var __1 = require("..");
var types_1 = require("../types");
function onMessage(socket) {
    socket.on(types_1.Events.MESSAGE, function (data) {
        if (data.message && typeof data.message == 'string') {
            __1.harmonyServer.emit('MESSAGE', __assign(__assign({}, data), { icon: lodash_1.default.get(__1.harmonyServer.getUsers()[socket.id], 'icon', '') }));
        }
    });
}
exports.default = onMessage;
