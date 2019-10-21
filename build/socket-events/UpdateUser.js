"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var jwt_1 = require("../promisified/jwt");
var __1 = require("..");
var userSchema_1 = require("../schema/userSchema");
function onUpdateUser(socket) {
    socket.on(types_1.Events.PROFILE_UPDATE, function (data) {
        if (typeof data.token === 'string') {
            jwt_1.verify(data.token, __1.config.config.jwtsecret)
                .then(function (result) {
                if (result.valid) {
                    if (result.decoded.userid) {
                        userSchema_1.User.findOne({ userid: result.decoded.userid })
                            .then(function (user) {
                            if (user) {
                                if (data.avatar) {
                                    user.avatar = data.avatar;
                                }
                                if (data.username) {
                                    user.username = data.username;
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
                    }
                }
            })
                .catch(function () {
                socket.emit(types_1.Events.PROFILE_UPDATE_ERROR, 'Missing Token');
            });
        }
    });
}
exports.default = onUpdateUser;
