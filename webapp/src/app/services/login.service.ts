import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import type { Observable } from 'rxjs';
import type { ILoginResponse } from '../interfaces/responses/login-response-interface';

@Injectable({
    providedIn: 'root',
})
export class LoginService {
    private readonly _httpClient = inject(HttpClient);

    login(email: string, password: string): Observable<ILoginResponse> {
        return this._httpClient.post<ILoginResponse>(
            'http://localhost:5000/login',
            {
                email,
                password,
            }
        );
    }
}
