import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { AngularMaterialModule } from './angular-material/angular-material.module';

@Component({
    selector: 'app-root',
    imports: [RouterOutlet, AngularMaterialModule],
    templateUrl: './app.component.html',
    styleUrl: './app.component.scss',
})
export class AppComponent {}
