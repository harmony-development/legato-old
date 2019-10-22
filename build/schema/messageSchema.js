"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var mongoose_1 = require("mongoose");
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
    }
});
exports.Message = mongoose_1.model('Message', exports.messageSchema);
