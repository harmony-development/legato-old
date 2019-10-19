"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var __1 = require("..");
function onProfileUpdate(socket) {
    socket.on(types_1.Events.PROFILE_UPDATE, function (data) {
        if (data.name && typeof data.name === 'string') {
            __1.harmonyServer.getUsers()[socket.id].name = data.name.substring(0, 50);
        }
        if (data.icon && typeof data.icon === 'string') {
            __1.harmonyServer.getUsers()[socket.id].icon = data.icon;
        }
    });
}
exports.default = onProfileUpdate;
