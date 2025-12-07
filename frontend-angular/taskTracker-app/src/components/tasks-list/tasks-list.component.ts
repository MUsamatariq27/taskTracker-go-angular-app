import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../services/api-service/api.service';
import { Task } from '../../models/task.model';
import { CommonModule,  } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-tasks-list',
  imports: [CommonModule, FormsModule],
  templateUrl: './tasks-list.component.html',
  styleUrl: './tasks-list.component.css'
})
export class TasksListComponent implements OnInit {

  tasks: Task[] = [];
  loading = true;
  error = '';
  showCreateForm = false;

  newTask = {
    title: '',
    description: '',
    completed: false, 
    user_id: null
  };

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    this.loadTasks();
  }

  loadTasks(): void {
    this.loading = true;
    const retrievedObject = JSON.parse(this.apiService.getUser());
    this.apiService.getUserTask(retrievedObject.id).subscribe({
      next: (res) => {
        this.tasks = res.tasks;
        this.loading = false;
      },
      error: (err) => {
        this.error = 'Failed to load tasks.';
        console.error(err);
        this.tasks = []
        this.loading = false;
      },
    });
  }

  toggleComplete(task: Task): void {
    task.completed = !task.completed;
    this.apiService.updateTask(task.id, task).subscribe({
      next: () => console.log('Task updated'),
      error: (err) => console.error(err),
    });
  }

editTask(task: Task): void {
  task.isEditing = true;
}

saveTask(task: Task): void {
  task.isEditing = false;
  this.apiService.updateTask(task.id, task).subscribe({
    next: () => console.log('Task updated'),
    error: (err) => console.error(err),
  });
}

deleteTask(task: Task){
  this.apiService.deleteTask(task.id).subscribe({
    next: () => {
      this.tasks = this.tasks.filter(ts => ts.id !== task.id);
    },
    error: (err) => console.error(err),
  });
}

  onCreateTask() {
    this.showCreateForm = true;
  }

  closeCreateTaskForm(){
    this.showCreateForm = false;
    this.newTask = {
      title: '',
      description: '',
      completed: false, 
      user_id: null
    }
  }

  addTask(){
    //console.log("add button clicked!");
    const userID = JSON.parse(this.apiService.getUser()).id
    this.newTask.user_id = userID

    this.apiService.createTask(this.newTask).subscribe({
        next: () => {
          this.loadTasks()
        },
        error: (err) => console.error(err),
    });
    this.closeCreateTaskForm()
  }

  
 

}
