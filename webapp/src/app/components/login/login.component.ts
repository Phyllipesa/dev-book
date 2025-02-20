import { Router } from '@angular/router';
import { Component, inject } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

import { LoginService } from '../../services/login.service';

@Component({
    selector: 'app-login',
    standalone: true,
    imports: [ReactiveFormsModule],
    templateUrl: './login.component.html',
    styleUrl: './login.component.scss',
})
export class LoginComponent {
    loginForm: FormGroup = new FormGroup({
        email: new FormControl(''),
        password: new FormControl(''),
    });

    private readonly _router = inject(Router);
    private readonly _loginService = inject(LoginService);

    onLogin(): void {
        this._loginService
            .login(this.loginForm.value.email, this.loginForm.value.password)
            .subscribe({
                next: (tokenResponse) => {
                    console.log('TOKEN', tokenResponse);
                    this._router.navigate(['criar-usuario']);
                },
                error: () => {},
            });
    }
}
