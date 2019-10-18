"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var __1 = require("..");
function onLogin(socket) {
    socket.on(types_1.Events.LOGIN, function (data) {
        if (data.name) {
            if (__1.harmonyServer.getUsers()[socket.id]) {
                __1.harmonyServer.getUsers()[socket.id].name = data.name;
                __1.harmonyServer.sendMessage(__1.harmonyServer.getUsers()[socket.id].name, 'has joined the channel');
            }
            else
                __1.harmonyServer.getUsers()[socket.id] = { name: data.name };
        }
    });
}
exports.default = onLogin;
