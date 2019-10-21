"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var isemail_1 = __importDefault(require("isemail"));
var types_1 = require("../types");
var __1 = require("..");
var jwt_1 = require("../promisified/jwt");
var userSchema_1 = require("../schema/userSchema");
function validPassword(password) {
    if (password.length < 5) {
        return { valid: false, message: 'Password must be at least 5 characters.' };
    }
    if (password.length > 30) {
        return {
            valid: false,
            message: 'Password cannot be longer than 30 characters.'
        };
    }
    if (!/[a-z]/.test(password)) {
        return { valid: false, message: 'Password must contain lowercase letters' };
    }
    if (!/[0-9]/.test(password)) {
        return { valid: false, message: 'Password must contain numbers' };
    }
    return { valid: true };
}
function onRegister(socket) {
    socket.on(types_1.Events.REGISTER, function (data) {
        if (typeof data.email === 'string' &&
            typeof data.password === 'string' &&
            typeof data.username === 'string') {
            if (isemail_1.default.validate(data.email, { errorLevel: true }) <= 5) {
                if (validPassword(data.password).valid) {
                    userSchema_1.User.findOne({ email: data.email }, function (err, user) {
                        if (!user) {
                            __1.harmonyServer.Database.register(data.email, data.password, data.username)
                                .then(function (user) {
                                jwt_1.sign({
                                    userid: user.userid
                                }, __1.config.config.jwtsecret, { expiresIn: '7d' })
                                    .then(function (token) {
                                    socket.emit(types_1.Events.REGISTER, {
                                        token: token,
                                        theme: {},
                                        username: data.username,
                                        avatar: ''
                                    });
                                })
                                    .catch(function (err) {
                                    console.log(err);
                                    socket.emit(types_1.Events.REGISTER_ERROR, 'Sorry, but the API is having a stroke right now');
                                });
                            })
                                .catch(function (err) {
                                console.log(err);
                                socket.emit(types_1.Events.REGISTER_ERROR, 'Sorry, but the API is having a stroke right now');
                            });
                        }
                        else {
                            socket.emit(types_1.Events.REGISTER_ERROR, 'Email already registered');
                        }
                    }).catch(function () {
                        socket.emit(types_1.Events.REGISTER_ERROR, 'Sorry, but the API is having a stroke right now');
                    });
                }
                else {
                    socket.emit(types_1.Events.REGISTER_ERROR, validPassword(data.password).message);
                }
            }
            else {
                socket.emit(types_1.Events.REGISTER_ERROR, 'Invalid Email.');
            }
        }
        else {
            socket.emit(types_1.Events.REGISTER_ERROR, 'Missing email or password.');
        }
    });
}
exports.default = onRegister;
