create table tb_departamento (
  id_depto serial not null primary key,
  nome varchar(100) not null
)

create table tb_funcionario(
  id_func serial not null primary key,
  nome varchar(100) not null,
  salario numeric(10,2) not null,
  id_depto integer not null,

  constraint fk_departamento
    foreign key (id_depto) references tb_departamento(id_depto)
    on delete cascade
    on update cascade
)
