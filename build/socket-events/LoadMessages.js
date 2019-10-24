"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
function onLoadMessages(socket) {
    socket.on(types_1.Events.LOAD_MESSAGES, function () { });
}
exports.default = onLoadMessages;
