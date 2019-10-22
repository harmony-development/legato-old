"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var jwt_1 = require("../promisified/jwt");
var __1 = require("..");
var userSchema_1 = require("../schema/userSchema");
function onGetUserData(socket) {
    socket.on(types_1.Events.GET_USER_DATA, function (data) {
        if (data.token) {
            jwt_1.verify(data.token, __1.config.config.jwtsecret)
                .then(function (result) {
                if (result.valid && result.decoded) {
                    if (result.decoded.userid) {
                        userSchema_1.User.findOne({ userid: result.decoded.userid })
                            .then(function (user) {
                            if (user) {
                                __1.harmonyServer.getSocketServer().emit(types_1.Events.GET_USER_DATA, {
                                    username: user.username,
                                    avatar: user.avatar,
                                    theme: user.theme
                                });
                            }
                            else {
                                socket.emit(types_1.Events.PROFILE_UPDATE_ERROR, 'You do not exist');
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
                .catch(function () { });
        }
        else {
            socket.emit(types_1.Events.GET_USER_DATA_ERROR, 'Invalid token');
        }
    });
}
exports.default = onGetUserData;
