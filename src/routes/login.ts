import { Router } from 'express';

const routes = Router();

routes.get('/login', (req, res) => {
  res.json({ status: 'online' });
});

export default routes;
