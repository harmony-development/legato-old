"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var express_1 = require("express");
var routes = express_1.Router();
routes.get('/register', function (req, res) {
    res.json({ status: 'online' });
});
exports.default = routes;
