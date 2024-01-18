<?php

declare(strict_types=1);

/**
 * This file is part of Hyperf.
 *
 * @link     https://www.hyperf.io
 * @document https://hyperf.wiki
 * @contact  group@hyperf.io
 * @license  https://github.com/hyperf/hyperf/blob/master/LICENSE
 */

use App\Middleware\JwtToUser;
use App\Middleware\Performance;
use Hyperf\HttpServer\Router\Router;

Router::addRoute(['GET', 'POST', 'HEAD'], '/', [App\Controller\IndexController::class, 'index'], [
    'middleware' => [Performance::class]
]);
Router::addRoute(['GET', 'POST', 'HEAD'], '/hello', [App\Controller\IndexController::class, 'hello'], [
    'middleware' => [Performance::class]
]);
Router::addRoute(['GET', 'POST', 'HEAD'], '/hello-world', [App\Controller\IndexController::class, 'helloWorld'], [
    'middleware' => [Performance::class]
]);
Router::addRoute(['GET', 'POST', 'HEAD'], '/hello-go', [App\Controller\IndexController::class, 'helloGo'], [
    'middleware' => [Performance::class]
]);


Router::addRoute('POST', '/login', App\Controller\LoginController::class, [
    'middleware' => [Performance::class]
]);

Router::addRoute(['GET', 'POST', 'HEAD'], '/try-check-jwt', [App\Controller\IndexController::class, 'tryAdd'], [
    'middleware' => [Performance::class, JwtToUser::class],
]);

Router::addRoute(['POST'], '/player-queue', [App\Controller\PlayerQueue::class, 'queue'], [
    'middleware' => [Performance::class],
]);
Router::addRoute(['POST'], '/player-queue-kafka', [App\Controller\PlayerQueue::class, 'kafkaQueue'], [
    'middleware' => [Performance::class],
]);


Router::get('/favicon.ico', function () {
    return '';
});

Router::addServer('grpc', function () {
    // Correspondence between the definition in the .proto file and
    // gRPC server routing: /{package}.{service}/{rpc}
    Router::addGroup('/hello.hi', function () {
        Router::post('/sayHello', [App\Controller\HiController::class, 'sayHello']);
    });
});