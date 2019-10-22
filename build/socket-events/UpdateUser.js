"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var __1 = require("..");
var userSchema_1 = require("../schema/userSchema");
function onUpdateUser(socket) {
    socket.on(types_1.Events.PROFILE_UPDATE, function (data) {
        __1.harmonyServer.Database.verifyToken(data.token)
            .then(function (userid) {
            userSchema_1.User.findOne({ userid: userid })
                .then(function (user) {
                if (user) {
                    if (data.avatar) {
                        user.avatar = data.avatar;
                    }
                    if (data.username) {
                        user.username = data.username;
                    }
                    if (data.theme) {
                        user.theme = data.theme;
                    }
                    user.save();
                }
                else {
                    socket.emit(types_1.Events.PROFILE_UPDATE_ERROR, 'You do not exist');
                }
            })
                .catch(function (err) {
                console.log(err);
                socket.emit(types_1.Events.PROFILE_UPDATE_ERROR, 'Failed to update profile');
            });
        })
            .catch(function () {
            __1.harmonyServer.SocketServer.emit(types_1.Events.INVALIDATE_SESSION, 'Invalid token');
        });
    });
}
exports.default = onUpdateUser;
