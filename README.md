Генерация отчета на основании сведений о продажах за последний месяц


GET /products - list of products -- 200, 404, 500


GET /products/:id - product by id -- 200, 404, 500


POST /products/:id - create product -- 204, 4xx


PUT /products/:id - fully update product -- 204/200, 400, 500


PATCH /products/:id - partially update product -- 204/200, 400, 500


DELETE /products/:id - delete product by id -- 204, 404, 400

