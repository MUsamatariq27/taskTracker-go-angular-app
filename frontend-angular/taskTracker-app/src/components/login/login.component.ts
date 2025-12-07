import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { RouterLink, Router } from '@angular/router';
import { ApiService } from '../../services/api-service/api.service';
import { OnInit } from '@angular/core';

@Component({
  selector: 'app-login',
  imports: [CommonModule, FormsModule, RouterLink],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent implements OnInit {

  email = '';
  password = '';
  errorMessage = '';

  constructor(private apiService: ApiService) {}
  router = inject(Router)

  ngOnInit(): void {
      if(this.apiService.isLoggedIn()){
        this.router.navigate(['/tasks']);
      }
  }

  

  onSubmit() {
    this.errorMessage = ''
    this.apiService.loginUser({email: this.email, password: this.password}).subscribe({
      next: (response) => {
        this.apiService.saveToken(response.token)
        let user = response.user
        //console.log(user.id)
        this.apiService.saveUser({email: user.email, id: user.id, name: user.name})
        this.router.navigate(['/tasks']);
      },
      error: (err) => {
        this.errorMessage = err.error?.error + " Please try Again."
      }
    });
  }

  onInputChange(){
    this.errorMessage = ''
  }

  

}
