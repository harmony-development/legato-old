"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var __1 = require("..");
function onUsernameUpdate(socket) {
    socket.on(types_1.Events.USERNAME_UPDATE, function (data) {
        if (data.name && typeof data.name == 'string') {
            __1.harmonyServer.sendMessage(__1.harmonyServer.getUsers()[socket.id].name || socket.id, "updated their username to " + data.name.substring(0, 50));
            __1.harmonyServer.getUsers()[socket.id].name = data.name.substring(0, 50);
        }
    });
}
exports.default = onUsernameUpdate;
