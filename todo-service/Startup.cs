using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.HttpsPolicy;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using TodoService.Repository;
using TodoService.Repository.InMemory;
using TodoService.Repository.Mongo;
using TodoService.Services;

namespace TodoService
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            services.AddMvc().SetCompatibilityVersion(CompatibilityVersion.Version_2_1);

            // Dependency Injection
            // services.Configure<TodoDatabaseSettings>(Configuration.GetSection(nameof(TodoDatabaseSettings)));
            // services.AddSingleton<ITodoDatabaseSettings>(sp => sp.GetRequiredService<IOptions<TodoDatabaseSettings>>().Value);
            // services.AddSingleton<MongoTodoRepository>();

            services.AddSingleton<InMemoryTodoDatabase>();
            services.AddSingleton<ITodoRepository>(sp => sp.GetRequiredService<InMemoryTodoDatabase>());

            services.AddSingleton<TodosService>();
            services.AddSingleton<ITodosService>(sp => sp.GetRequiredService<TodosService>());

        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IHostingEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }
            // else
            // {
            //     app.UseHsts();
            // }

            // app.UseHttpsRedirection();
            app.UseMvc();
        }
    }
}
