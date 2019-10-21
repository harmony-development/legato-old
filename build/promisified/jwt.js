"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var jsonwebtoken_1 = __importDefault(require("jsonwebtoken"));
function sign(payload, secretOrPrivateKey, options) {
    return new Promise(function (resolve, reject) {
        if (options) {
            jsonwebtoken_1.default.sign(payload, secretOrPrivateKey, options, function (err, token) {
                if (err) {
                    reject(err);
                    return;
                }
                resolve(token);
            });
        }
        else {
            jsonwebtoken_1.default.sign(payload, secretOrPrivateKey, function (err, token) {
                if (err) {
                    reject(err);
                    return;
                }
                resolve(token);
            });
        }
    });
}
exports.sign = sign;
function verify(token, secretOrPublicKey, options) {
    return new Promise(function (resolve, reject) {
        jsonwebtoken_1.default.verify(token, secretOrPublicKey, options, function (err, decoded) {
            if (decoded && !err) {
                resolve({ valid: true, decoded: decoded });
            }
            else {
                resolve({ valid: false });
            }
        });
    });
}
exports.verify = verify;
