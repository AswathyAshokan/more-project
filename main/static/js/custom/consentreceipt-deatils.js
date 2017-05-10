 $(document).ready(function() {
     
     var table = $('#tbl_details').DataTable( {
                   "order": [[1, 'asc']]
        } );

        // Add event listener for opening and closing details
        $('#tbl_details tbody').on('click', 'td.details-control', function () {
            var tr = $(this).closest('tr');
            var row = table.row( tr );

            if ( row.child.isShown() ) {
                // This row is already open - close it
                row.child.hide();
                tr.removeClass('shown');
            }
            else {
                // Open this row
                row.child( format(row.data()) ).show();
                tr.addClass('shown');
            }
        } );
    } );