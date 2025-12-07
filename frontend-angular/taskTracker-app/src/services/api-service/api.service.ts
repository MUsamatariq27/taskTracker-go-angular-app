import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Task } from '../../models/task.model';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  private baseUrl = 'http://localhost:8080/api'; // Go backend URL

  constructor(private http: HttpClient) {}

  // Auth
  registerUser(user: any): Observable<any> {
    return this.http.post(`${this.baseUrl}/auth/register`, user);
  }

  loginUser(credentials: any): Observable<any> {
    return this.http.post(`${this.baseUrl}/auth/login`, credentials);
  }

   getToken(): string | null {
    return localStorage.getItem('token');
  }

  isLoggedIn(): boolean {
    return !!this.getToken();
  }

  logout() {
    localStorage.removeItem('token');
    localStorage.removeItem('userInfo');
  }

  saveToken(token: string) {
    localStorage.setItem('token', token);
  }

  saveUser(userInfo: any) {
    localStorage.setItem('userInfo', JSON.stringify(userInfo));
  }

  getUser(): any | null {
    return localStorage.getItem('userInfo');
  }

  
  getAllTasks(): Observable<any> {
    return this.http.get(`${this.baseUrl}/tasks`);
  }

  createTask(task: any): Observable<any> {
    return this.http.post(`${this.baseUrl}/task`, JSON.parse(JSON.stringify(task)));
  }

  updateTask(id: number, task: any): Observable<any> {
    return this.http.put(`${this.baseUrl}/task/${id}`, task);
  }

  deleteTask(id: number): Observable<any> {
    return this.http.delete(`${this.baseUrl}/task/${id}`);
  }

  getUserTask(id: number){
     return this.http.get<{ tasks: Task[] }>(`${this.baseUrl}/tasks/${id}`);
  }
}
