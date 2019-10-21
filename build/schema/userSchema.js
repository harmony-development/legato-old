"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var mongoose_1 = __importDefault(require("mongoose"));
var bcrypt_1 = __importDefault(require("bcrypt"));
var crypto_random_string_1 = __importDefault(require("crypto-random-string"));
exports.userSchema = new mongoose_1.default.Schema({
    userid: {
        unique: true,
        required: false,
        type: String
    },
    username: {
        unique: false,
        required: true,
        type: String
    },
    email: {
        unique: false,
        required: true,
        type: String
    },
    password: {
        unique: false,
        required: true,
        type: String
    },
    theme: {
        unique: false,
        required: false,
        type: {
            primary: {
                light: String,
                dark: String,
                contrastText: String
            },
            secondary: {
                light: String,
                dark: String,
                contrastText: String
            },
            type: String
        }
    }
});
exports.userSchema.pre('save', function (next) {
    var _this = this;
    bcrypt_1.default
        .hash(this.password, 10)
        .then(function (hash) {
        _this.password = hash;
        _this.userid = crypto_random_string_1.default({ length: 15 });
        next();
    })
        .catch(function (err) {
        next(err);
    });
});
exports.User = mongoose_1.default.model('User', exports.userSchema);
