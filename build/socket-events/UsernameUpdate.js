"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var lodash_1 = __importDefault(require("lodash"));
var types_1 = require("../types");
var __1 = require("..");
function onUsernameUpdate(socket) {
    socket.on(types_1.Events.USERNAME_UPDATE, function (data) {
        if (data.name && typeof data.name == 'string') {
            __1.harmonyServer.emit('MESSAGE', {
                author: lodash_1.default.get(__1.harmonyServer.getUsers()[socket.id], 'name', socket.id),
                message: "has changed their name to " + data.name.substring(0, 50)
            });
            __1.harmonyServer.getUsers()[socket.id].name = data.name.substring(0, 50);
        }
    });
}
exports.default = onUsernameUpdate;
