<?php

declare(strict_types=1);


namespace App\Middleware;

use App\Services\Jwt\JwtService;
use Hyperf\HttpServer\Response;
use Psr\Container\ContainerInterface;
use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\MiddlewareInterface;
use Psr\Http\Server\RequestHandlerInterface;

class JwtToUser implements MiddlewareInterface
{
    private const BEARER_REGEX = '/Bearer\s+(.*)$/i';
    /**
     * @var JwtService
     */
    private mixed $jwt;

    public function __construct(protected ContainerInterface $container)
    {
        $this->jwt = $this->container->get(JwtService::class);
    }

    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        $auth = $request->getHeaderLine('Authorization');
        if (!preg_match(self::BEARER_REGEX, $auth, $matches)) {
            return (new Response())->json([
                'message' => 'Authorization header is missing'
            ])->withStatus(403);
        }

        $token = $matches[1];
        try {
            $result = $this->jwt->decodeToken($token);
            $request = $request->withAttribute('user', $result);
            return $handler->handle($request);
        } catch (\LogicException $e) {
            // errors having to do with environmental setup or malformed JWT Keys
            return (new Response())->json([
                'message' => $e->getMessage()
            ])->withStatus(403);
        } catch (\UnexpectedValueException $e) {
            return (new Response())->json([
                'message' => $e->getMessage()
            ])->withStatus(403);
        }
    }

}