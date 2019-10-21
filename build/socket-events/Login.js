"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var types_1 = require("../types");
var bcrypt_1 = __importDefault(require("bcrypt"));
var userSchema_1 = require("../schema/userSchema");
var __1 = require("..");
var jwt_1 = require("../promisified/jwt");
function onLogin(socket) {
    socket.on(types_1.Events.LOGIN, function (data) {
        if (data.email && data.password) {
            userSchema_1.User.findOne({ email: data.email })
                .then(function (user) {
                if (user) {
                    if (user.password) {
                        bcrypt_1.default
                            .compare(data.password, user.password)
                            .then(function (success) {
                            if (success) {
                                jwt_1.sign({
                                    userid: user.userid
                                }, __1.config.config.jwtsecret, { expiresIn: '7d' })
                                    .then(function (token) {
                                    socket.emit(types_1.Events.LOGIN, {
                                        token: token,
                                        theme: user.theme,
                                        username: user.username,
                                        avatar: user.avatar
                                    });
                                })
                                    .catch(function () {
                                    socket.emit(types_1.Events.LOGIN_ERROR, 'Uhm. The API is having a stroke.');
                                });
                            }
                            else {
                                socket.emit(types_1.Events.LOGIN_ERROR, 'Invalid email or password');
                            }
                        })
                            .catch(function () {
                            socket.emit(types_1.Events.LOGIN_ERROR, 'Uhm. The API is having a stroke.');
                        });
                    }
                    else {
                        socket.emit(types_1.Events.LOGIN_ERROR, 'Somehow the password is missing in our records. Email support please!');
                    }
                }
                else {
                    socket.emit(types_1.Events.LOGIN_ERROR, 'Invalid email or password');
                }
            })
                .catch(function (err) {
                console.log(err);
                socket.emit(types_1.Events.LOGIN_ERROR, 'Invalid email or password');
            });
        }
        else {
            socket.emit(types_1.Events.LOGIN_ERROR, 'Missing username or password');
        }
    });
}
exports.default = onLogin;
