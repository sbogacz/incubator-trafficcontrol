use utf8;
package Schema::Result::FederationFederationResolver;

# Created by DBIx::Class::Schema::Loader
# DO NOT MODIFY THE FIRST PART OF THIS FILE

=head1 NAME

Schema::Result::FederationFederationResolver

=cut

use strict;
use warnings;

use base 'DBIx::Class::Core';

=head1 TABLE: C<federation_federation_resolver>

=cut

__PACKAGE__->table("federation_federation_resolver");

=head1 ACCESSORS

=head2 federation

  data_type: 'bigint'
  is_foreign_key: 1
  is_nullable: 0

=head2 federation_resolver

  data_type: 'bigint'
  is_foreign_key: 1
  is_nullable: 0

=head2 last_updated

  data_type: 'timestamp with time zone'
  default_value: current_timestamp
  is_nullable: 1
  original: {default_value => \"now()"}

=cut

__PACKAGE__->add_columns(
  "federation",
  { data_type => "bigint", is_foreign_key => 1, is_nullable => 0 },
  "federation_resolver",
  { data_type => "bigint", is_foreign_key => 1, is_nullable => 0 },
  "last_updated",
  {
    data_type     => "timestamp with time zone",
    default_value => \"current_timestamp",
    is_nullable   => 1,
    original      => { default_value => \"now()" },
  },
);

=head1 PRIMARY KEY

=over 4

=item * L</federation>

=item * L</federation_resolver>

=back

=cut

__PACKAGE__->set_primary_key("federation", "federation_resolver");

=head1 RELATIONS

=head2 federation

Type: belongs_to

Related object: L<Schema::Result::Federation>

=cut

__PACKAGE__->belongs_to(
  "federation",
  "Schema::Result::Federation",
  { id => "federation" },
  { is_deferrable => 0, on_delete => "CASCADE", on_update => "CASCADE" },
);

=head2 federation_resolver

Type: belongs_to

Related object: L<Schema::Result::FederationResolver>

=cut

__PACKAGE__->belongs_to(
  "federation_resolver",
  "Schema::Result::FederationResolver",
  { id => "federation_resolver" },
  { is_deferrable => 0, on_delete => "CASCADE", on_update => "CASCADE" },
);


# Created by DBIx::Class::Schema::Loader v0.07045 @ 2016-11-15 09:35:47
# DO NOT MODIFY THIS OR ANYTHING ABOVE! md5sum:+rDQ18D+CIYbqA/JYyr4UA


# You can replace this text with custom code or comments, and it will be preserved on regeneration
1;
