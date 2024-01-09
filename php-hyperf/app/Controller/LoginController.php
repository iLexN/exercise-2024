<?php

declare(strict_types=1);


namespace App\Controller;

use App\Services\Jwt\UserGenJwt;
use App\Services\Jwt\JwtService;
use App\Services\Login\LoginForm;

class LoginController extends AbstractController
{
    public function __construct(
        private readonly JwtService $userJwt,
    )
    {
    }

    public function __invoke()
    {
        try {
            $login = LoginForm::fromRequest($this->request);
        } catch (\Throwable) {
            return $this->response->json([
                'error' => "Invalid request"
            ])->withStatus(400);
        }

        $userGenJwt = UserGenJwt::createFromLoginForm($login);

        return $this->response->json([
            'message' => "Login successful",
            'access_token' => $this->userJwt->genToken($userGenJwt),
        ]);
    }
}