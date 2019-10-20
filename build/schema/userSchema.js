"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var mongoose_1 = __importDefault(require("mongoose"));
var bcrypt_1 = __importDefault(require("bcrypt"));
exports.userSchema = new mongoose_1.default.Schema({
    userid: {
        unique: true,
        required: true,
        type: String
    },
    username: {
        unique: false,
        required: true,
        type: String
    },
    password: {
        unique: false,
        required: true,
        type: String
    }
});
exports.userSchema.pre('save', function (next) {
    var _this = this;
    bcrypt_1.default.hash(this.password, 10, function (err, hash) {
        if (err) {
            return next(err);
        }
        _this.password = hash;
        next();
    });
});
exports.User = mongoose_1.default.model('User', exports.userSchema);
