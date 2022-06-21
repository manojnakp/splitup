import { Router } from 'express';
const router = Router();

/* GET users listing. */
router.get('/', function(req, res) {
  res.end('respond with a resource');
});

export default router;
