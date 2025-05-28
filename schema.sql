create table if not exists tbl_professor (
	func 	    serial not null primary key,
	nome 	    varchar(100) not null,
  email     varchar(100) not null unique,
  telefone  varchar(15) not null,
)

create table if not exists tbl_disciplina (
  codigo 	  serial not null primary key,
  nome 		  varchar(100) not null,
  ch        integer not null,
  ementa    text,
  biblio    text,
  conteudo  text,
  just      text,
  prof_resp integer,
  
  constraint fk_professor
    foreign key (prof_resp) references tbl_professor(func)
    on delete set null
    on update no action
)

create table if not exists tbl_curso (
  cod           serial not null primary key,
  nome          varchar(100) not null,
  duracao       integer not null,
  autorizacao   varchar(100) not null,
  prof_coord    integer,

  constraint fk_professor_coord
    foreign key (prof_coord) references tbl_professor(func)
    on delete set null
    on update no action
)

create table if not exists tbl_grade (
  numSeq      serial not null primary key,
  anoInicio   integer not null,
  semInicio  integer not null,
  codCurso    integer,

  constraint fk_curso
    foreign key (codCurso) references tbl_curso(cod)
    on delete set null
    on update no action
)

create table if not exists tbl_grade_disc (
  numSeq integer not null,
  codigo  integer not null,

  constraint fk_grade
    foreign key (numSeq) references tbl_grade(numSeq)
    on delete cascade
    on update no action,

  constraint fk_disciplina
    foreign key (codigo) references tbl_disciplina(codigo)
    on delete cascade
    on update no action,

  constraint pk_grade_disc primary key (numSeq, codigo)
)

create table if not exists tbl_ministra (
  func integer not null,
  codigo integer not null,
  ano_ministra integer not null,
  sem_ministra integer not null,

  constraint fk_professor_ministra
    foreign key (func) references tbl_professor(func)
    on delete cascade
    on update no action,

  constraint fk_disciplina_ministra
    foreign key (codigo) references tbl_disciplina(codigo)
    on delete cascade
    on update no action,

 constraint pk_tbl_ministra primary key (func, codigo, ano_ministra, sem_ministra)
)