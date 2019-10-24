"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var mongoose_1 = require("mongoose");
var crypto_random_string_1 = __importDefault(require("crypto-random-string"));
exports.messageSchema = new mongoose_1.Schema({
    author: {
        unique: false,
        required: true,
        type: String
    },
    message: {
        unique: false,
        required: true,
        type: String
    },
    files: {
        unique: false,
        required: true,
        type: Array()
    },
    messageid: {
        unique: true,
        required: true,
        type: String
    }
}, { timestamps: { createdAt: 'created_at' } });
exports.messageSchema.pre('validate', function (next) {
    //if (!this.isModified('messageid')) return next();
    this.messageid = crypto_random_string_1.default({ length: 30 });
    next();
});
exports.Message = mongoose_1.model('Message', exports.messageSchema);
