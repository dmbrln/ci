/* @ts-check */

export default (next, request, response, app) => next(
    console.log(request, response, app)
);
