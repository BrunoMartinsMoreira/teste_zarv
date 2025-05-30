insert into tb_departamento (nome) 
values ('Recursos Humanos'), ('Financeiro'),  ('TI');

insert into tb_funcionario (nome, salario, id_depto) 
values 
  ('João Silva', 3000.00, 1);
  ('Maria Oliveira', 3500.00, 2);
  ('Carlos Pereira', 4000.00, 3);

with menor as (
  select 
    id_depto,
    min(salario) as menor_salario
  from tb_funcionario
  group by id_depto
),
maior as (
  select 
    id_depto,
    max(salario) as maior_salario
  from tb_funcionario
  group by id_depto
)
select 
  td.nome as departamento,
  round(avg(tf.salario), 2) as media_salario,
  menor.menor_salario,
  maior.maior_salario,
  min(case when tf.salario = menor.menor_salario then tf.nome end) as funcionario_menor_salario,
  min(case when tf.salario = maior.maior_salario then tf.nome end) as funcionario_maior_salario
from tb_funcionario tf
left join tb_departamento td on tf.id_depto = td.id_depto
left join menor on tf.id_depto = menor.id_depto
left join maior on tf.id_depto = maior.id_depto
group by td.nome, menor.menor_salario, maior.maior_salario