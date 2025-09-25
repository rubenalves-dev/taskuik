import { AsyncPipe } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { Component, effect, inject, OnInit, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Task, TasksService } from './service/tasks.service';
import { Observable } from 'rxjs';
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';

@Component({
    selector: 'app-root',
    imports: [RouterOutlet, AsyncPipe, ReactiveFormsModule],
    templateUrl: './app.html',
    styleUrl: './app.scss',
})
export class App implements OnInit {
    protected readonly tasksService = inject(TasksService);

    protected readonly title = signal('client');

    protected tasks$!: Observable<Task[]>;

    protected taskForm = new FormControl('');

    ngOnInit(): void {
        this.refreshTasks();

        this.tasks$.subscribe((tasks) => {
            console.log('Tasks updated:', tasks);
        });
    }

    public onSubmit() {
        const title = this.taskForm.get('title')?.value;
        this.tasksService.addTask(title || '').subscribe(() => {
            this.taskForm.reset();
            this.refreshTasks();
        });
    }

    private refreshTasks() {
        this.tasks$ = this.tasksService.getTasks();
    }
}
