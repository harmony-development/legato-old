"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var __1 = require("..");
function onLogin(socket) {
    socket.on(types_1.Events.LOGIN, function (data) {
        if (data.name) {
            if (__1.harmonyServer.getUsers()[socket.id]) {
                __1.harmonyServer.emit('MESSAGE', {
                    author: __1.harmonyServer.getUsers()[socket.id].name,
                    message: "has joined the channel"
                });
                __1.harmonyServer.getUsers()[socket.id].name = data.name;
            }
            else
                __1.harmonyServer.getUsers()[socket.id] = { name: data.name };
        }
    });
}
exports.default = onLogin;
