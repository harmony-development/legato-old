"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var jwt_1 = require("../promisified/jwt");
var __1 = require("..");
var userSchema_1 = require("../schema/userSchema");
function onMessage(socket) {
    socket.on(types_1.Events.MESSAGE, function (data) {
        jwt_1.verify(data.token, __1.config.config.jwtsecret)
            .then(function (result) {
            if (result.valid && result.decoded) {
                if (result.decoded.userid) {
                    userSchema_1.User.findOne({ userid: result.decoded.userid })
                        .then(function (user) {
                        if (user) {
                            __1.harmonyServer.getSocketServer().emit(types_1.Events.MESSAGE, {
                                username: user.username,
                                message: data.message,
                                avatar: user.avatar,
                                files: data.files
                            });
                        }
                    })
                        .catch(function (err) {
                        console.log(err);
                    });
                }
                else {
                    socket.emit(types_1.Events.INVALIDATE_SESSION, 'invalid session token');
                }
            }
            else {
                socket.emit(types_1.Events.INVALIDATE_SESSION, 'invalid session token');
            }
        })
            .catch(function () {
            socket.emit(types_1.Events.INVALIDATE_SESSION, 'invalid session token');
        });
    });
}
exports.default = onMessage;
