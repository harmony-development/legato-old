"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var __1 = require("..");
var userSchema_1 = require("../schema/userSchema");
function onGetUserData(socket) {
    socket.on(types_1.Events.GET_USER_DATA, function (data) {
        __1.harmonyServer.Database.verifyToken(data.token)
            .then(function (userid) {
            if (data.targetuser) {
                userSchema_1.User.findOne({ userid: data.targetuser })
                    .then(function (user) {
                    if (user) {
                        __1.harmonyServer
                            .getSocketServer()
                            .emit(types_1.Events.GET_TARGET_USER_DATA, {
                            userid: user.userid,
                            username: user.username,
                            avatar: user.avatar
                        });
                    }
                })
                    .catch(function (err) {
                    console.log(err);
                });
            }
            else {
                userSchema_1.User.findOne({ userid: userid })
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
        })
            .catch(function () {
            socket.emit(types_1.Events.INVALIDATE_SESSION, 'Invalid token');
        });
    });
}
exports.default = onGetUserData;
