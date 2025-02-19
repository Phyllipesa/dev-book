import { Component, inject } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';

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

    onLogin(): void {}
}
