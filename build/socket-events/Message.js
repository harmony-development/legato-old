"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var __1 = require("..");
var types_1 = require("../types");
function onMessage(socket) {
    socket.on(types_1.Events.MESSAGE, function (data) {
        if (data.message && typeof data.message == 'string') {
            if (__1.harmonyServer.getUsers()[socket.id]) {
                __1.harmonyServer.sendMessage(__1.harmonyServer.getUsers()[socket.id].name || socket.id, data.message.substring(0, 500));
            }
            else {
                __1.harmonyServer.sendMessage(socket.id, data.message.substring(0, 500));
            }
        }
    });
}
exports.default = onMessage;
