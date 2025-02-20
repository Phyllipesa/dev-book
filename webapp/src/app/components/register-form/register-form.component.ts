import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

@Component({
    selector: 'app-register-form',
    standalone: true,
    imports: [ReactiveFormsModule],
    templateUrl: './register-form.component.html',
    styleUrl: './register-form.component.scss',
})
export class RegisterFormComponent {
    registerForm: FormGroup = new FormGroup({
        name: new FormControl(''),
        nick: new FormControl(''),
        email: new FormControl(''),
        password: new FormControl(''),
    });
}
