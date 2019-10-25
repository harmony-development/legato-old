"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var __1 = require("..");
var userSchema_1 = require("../schema/userSchema");
function onMessage(socket) {
    socket.on(types_1.Events.MESSAGE, function (data) {
        __1.harmonyServer.Database.verifyToken(data.token)
            .then(function (userid) {
            userSchema_1.User.findOne({ userid: userid })
                .then(function (user) {
                if (user) {
                    __1.harmonyServer.Database.addMessage(userid, data.message, data.files).then(function () {
                        __1.harmonyServer.getSocketServer().emit(types_1.Events.MESSAGE, {
                            author: userid,
                            avatar: user.avatar,
                            message: data.message,
                            files: data.files
                        });
                    });
                }
            })
                .catch(function (err) {
                console.log(err);
            });
        })
            .catch(function () {
            __1.harmonyServer
                .getSocketServer()
                .emit(types_1.Events.INVALIDATE_SESSION, 'Invalid token');
        });
    });
}
exports.default = onMessage;
