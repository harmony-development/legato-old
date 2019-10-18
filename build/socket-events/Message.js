"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var __1 = require("..");
var types_1 = require("../types");
function onMessage(socket) {
    socket.on(types_1.Events.MESSAGE, function (data) {
        if (data.message && typeof data.message == 'string') {
            __1.harmonyServer.emit('MESSAGE', data);
        }
    });
}
exports.default = onMessage;
