package com.example.Test.Dashboard;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import org.springframework.cloud.netflix.hystrix.dashboard.EnableHystrixDashboard;

@EnableHystrixDashboard

@SpringBootApplication
public class TestDashboardApplication {

	public static void main(String[] args) {
		SpringApplication.run(TestDashboardApplication.class, args);
	}
}
