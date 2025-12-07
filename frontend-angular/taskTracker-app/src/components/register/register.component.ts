import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterLink, Router } from '@angular/router';
import { ApiService } from '../../services/api-service/api.service';

@Component({
  selector: 'app-register',
  imports: [CommonModule, FormsModule, RouterLink],
  templateUrl: './register.component.html',
  styleUrl: './register.component.css'
})
export class RegisterComponent {

  name = '';
  email = '';
  password = '';
  error = '';

  constructor(private apiService: ApiService) {}
  router = inject(Router)

  onSubmit() {
    this.apiService.registerUser({name: this.name, email: this.email, password: this.password}).subscribe({
      next: (response) => {
        console.log(response["message"])
        alert(response["message"] + ". Please LogIn!")
        this.router.navigate(['/login'])
      },
      error: (err) => {
        if (err.status === 400 && err.error?.error?.includes('Exists')) {
          this.error = 'This email is already registered. Please try logging in.';
          alert(this.error)
          this.router.navigate(['/login'])
        } else {
          alert('Registration Unsuccesful. Please try again.');
        }
      }
    });
  }

}
